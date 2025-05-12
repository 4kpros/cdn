# Overview

This project leverages Docker and Make to streamline development and deployment.

# Prerequisites

- Docker installed and running

- Make installed

- A GitHub account with a personal access token with the necessary permissions

# Usage

1. Build and run docker image

    To build and run locally using Docker, run the following command:
    ```
        make docker-cdn
    ```
    This will build and start the Docker container.

2. Logging into GitHub Container Registry

    Before pushing or pulling images, you need to log in to the GitHub Container Registry:
    ```
        make docker-ghcr-login
    ```
    You will be prompted to enter your GitHub personal access token.

3. Pushing Image to GitHub Container Registry

    To push a new image to the GitHub Container Registry, run:
    ```
        make docker-ghcr-push
    ```
    You will be prompted to enter the new version tag.

4. Pulling Image from GitHub Container Registry

    To pull a specific image from the GitHub Container Registry, run:
    ```
        make docker-ghcr-pull
    ```
    You will be prompted to enter the desired image tag.


# Update GitHub Action Secrets for continuous integration(build and package)

Go to this link: [GitHub Action Secrets](https://github.com/EMENEC-FINANCE/cdn/settings/secrets/actions)

- ------------- On your GitHub Action Secrets page -------------
    - Set Secrets `GHCR_USERNAME` `GHCR_PASSWORD` with value your GitHub credentials. `GHCR_PASSWORD` is your personal access token with `write package` permission enabled
        

# Makefile Targets

- `docker-cdn`: Builds and starts the Docker container for local development.

- `docker-ghcr-push`: Builds the Docker image and pushes it to the GitHub Container Registry.

- `docker-ghcr-pull`: Pulls a specific image from the GitHub Container Registry.


# Additional Notes

By following these steps and customizing the Makefile to fit your specific needs, you can effectively manage your project using Docker and Make.