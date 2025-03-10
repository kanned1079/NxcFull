package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"runtime"
	"sync"
	"time"
)

// CacheEntry 包含缓存的 DNS 记录和过期时间
type CacheEntry struct {
	Answers   []dns.RR
	ExpiresAt time.Time
}

// 缓存结构
var cache = struct {
	sync.RWMutex
	data map[string]CacheEntry
}{data: make(map[string]CacheEntry)}

// resolver 用于从上游 DNS 服务器获取 DNS 记录并缓存
func resolver(domain string, qtype uint16) []dns.RR {
	// 先检查缓存中是否有结果
	cache.RLock() // 使用 RLock 读取缓存
	if cachedAnswer, found := cache.data[domain]; found {
		if time.Now().Before(cachedAnswer.ExpiresAt) {
			cache.RUnlock() // 正常解锁
			log.Printf("[CACHE HIT] : %s\n", domain)
			return cachedAnswer.Answers
		}
		// 如果缓存过期，删除该缓存
		cache.RUnlock() // 先解锁
		cache.Lock()    // 然后获取写锁
		delete(cache.data, domain)
		cache.Unlock() // 释放写锁
	}
	cache.RUnlock() // 之前是没有 RLock，所以要加这个

	// 如果没有缓存，或者缓存已过期，发送请求到上游 DNS 服务器
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), qtype)
	m.RecursionDesired = true

	c := &dns.Client{Timeout: 5 * time.Second}
	response, _, err := c.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Printf("[ERROR] : %v\n", err)
		return nil
	}

	if response == nil {
		log.Printf("[ERROR] : no response from server\n")
		return nil
	}

	// 将查询结果缓存，设置过期时间为 TTL
	cache.Lock() // 加写锁
	for _, answer := range response.Answer {
		if aRecord, ok := answer.(*dns.A); ok {
			// 假设 TTL 为 1 小时
			cache.data[domain] = CacheEntry{
				Answers:   response.Answer,
				ExpiresAt: time.Now().Add(time.Duration(aRecord.Hdr.Ttl) * time.Second),
			}
			break
		}
	}
	cache.Unlock() // 释放写锁

	log.Printf("[CACHE MISS] : %s\n", domain)
	for _, answer := range response.Answer {
		fmt.Printf("%s\n", answer.String())
	}

	return response.Answer
}

// dnsHandler 处理 DNS 查询请求
type dnsHandler struct{}

func (h *dnsHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		answers := resolver(question.Name, question.Qtype)
		if answers == nil {
			msg.SetRcode(r, dns.RcodeServerFailure)
		} else {
			msg.Answer = append(msg.Answer, answers...)
		}
	}

	w.WriteMsg(msg)
}

// StartDNSServer 启动 DNS 服务器
func StartDNSServer() {
	handler := new(dnsHandler)
	server := &dns.Server{
		Addr:      ":53",
		Net:       "udp",
		Handler:   handler,
		UDPSize:   65535,
		ReusePort: true,
	}

	tcpServer := &dns.Server{
		Addr:      ":53",
		Net:       "tcp",
		Handler:   handler,
		UDPSize:   65535,
		ReusePort: true,
	}

	fmt.Println("Starting DNS server on port 53 (UDP and TCP)")

	// 启动 UDP 服务器
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to start UDP server: %s\n", err.Error())
		}
	}()

	// 启动 TCP 服务器
	go func() {
		err := tcpServer.ListenAndServe()
		if err != nil {
			fmt.Printf("Failed to start TCP server: %s\n", err.Error())
		}
	}()

	select {} // 阻塞主 goroutine 直到服务器退出
}

func main() {
	go StartDNSServer()
	for {
		log.Println("当前偕程数 ", runtime.NumGoroutine())
		time.Sleep(time.Second * 5)
	}
}
