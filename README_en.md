## This is a Test Card Selling System

[Chinese](https://github.com/kanned1079/NxcFull/blob/main/README.md) | English

### Disclaimer

1. This software is independently developed by an individual and is intended to provide specific functions and services to users. Due to limited development resources and capabilities, the software may have some imperfections.
2. The individual developer assumes no responsibility for any consequences arising from the use of this software, including but not limited to device damage, data loss, profit loss, or business interruption.
3. The individual developer does not guarantee the accuracy, completeness, timeliness, or reliability of the information, functions, and services contained in the software. Users should verify and assess the information themselves and take responsibility for their own usage behavior.

### Building the Frontend Application

The frontend consists of two parts: `the main user and admin interface` and `a separate interface for configuring microservices`. Therefore, dependencies need to be installed and built in different directories.

#### Installing Frontend Project Dependencies

The project uses `npm` to manage dependencies.

```shell
make install || npm install # 安裝依賴
```

Run the build process:

```shell
# Requires runtime environment node (20.16+)
make build || npm run build # Use Vite for bundling
```

This will generate a dist build folder in `nxc-frontend/`. Next, refer to `nxc-frontend/INSTALL.md` for installation and execution instructions.


### Building the Backend Application

#### Unified Build

This is the recommended way to build the backend.

Run the following command in the project’s `backend/` directory to build all backend services:
```shell
# Requires docker (27.1+), golang (1.22.2+), and make installed on the system executing the build.
./build_all linux amd64 # Builds executable files and Docker images for the Linux amd64 architecture
```

After the build is complete, you will find all microservices’ Docker image packages `*.tar.gz` in the `backend/export/` directory, along with:
- `README.md` A guide for installation on the server environment.
- `docker-compose.yaml`  A configuration file for running services using Docker Compose.
- `manage_docker_images.sh` A script for quickly importing and deleting Docker images on the server.

At this point, all backend programs have been packaged. Next, refer to `backend/README.md` for installation and execution instructions.

#### Individual Service Build

All microservices are located in `./backend/*`.

To build an executable file, run `./run_build <OS> <Architecture>` in each microservice directory:

```shell
# Build executable files
./run_build.sh linux amd64 # Build for Linux amd64
./run_build.sh build-darwin-arm64 # Build for macOS arm64
./run_build.sh build-windows-amd64 # Build for Windows amd64
```

After the build is complete, the `build` directory within the current service will contain the corresponding executable file `exec and the Docker image.

