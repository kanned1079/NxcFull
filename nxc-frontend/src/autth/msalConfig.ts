// authConfig.ts
import { type Configuration } from "@azure/msal-browser";

export const msalConfig: Configuration = {
    auth: {
        clientId: "你的应用ID",  // Azure中获取的Application (client) ID
        authority: "https://login.microsoftonline.com/你的租户ID", // Azure中获取的Directory (tenant) ID
        redirectUri: "http://localhost:8080"  // 本地开发时的重定向URI
    }
};

export const loginRequest = {
    scopes: ["User.Read"]
};