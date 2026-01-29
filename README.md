# Overview

This project leverages Docker and Make to streamline development and deployment.

# Prerequisites

- Docker installed and running

- Make installed

# Pull and run docker image locally

1. Setup the environment

   Clone the repository:

   ```
       git clone https://github.com/4kpros/cdn.git
   ```

   ```
       cd cdn
   ```

   ```
    cp .env.example app.env
   ```

2. Run the CDN

   Pull the Docker image from GitHub Container Registry:

   ```
       make docker-cdn
   ```


# Build your own docker image

1. Clone, build and run docker image

   Clone the repository:

   ```
       git clone https://github.com/4kpros/cdn.git
   ```

   ```
       cd cdn
   ```

   To build and run locally using Docker, run the following command:

   ```
       make docker-cdn-build
   ```

   This will build and start the Docker container.

# Update GitHub Action Secrets for continuous integration(build and package)

Go to this link: [GitHub Action Secrets](https://github.com/4kpros/cdn/settings/secrets/actions)

- ------------- On your GitHub Action Secrets page -------------

  - Set Secrets `GHCR_USERNAME` `GHCR_PASSWORD` with value your GitHub credentials. `GHCR_PASSWORD` is your personal access token with `write package` permission enabled

# Makefile Targets

- `docker-cdn`: Builds and starts the Docker container for local development.

- `docker-ghcr-push`: Builds the Docker image and pushes it to the GitHub Container Registry.

- `docker-ghcr-pull`: Pulls a specific image from the GitHub Container Registry.

# Additional Notes

By following these steps and customizing the Makefile to fit your specific needs, you can effectively manage your project using Docker and Make.
