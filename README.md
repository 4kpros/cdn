# Customizable CDN with Go and Docker

- A flexible CDN built with Go, offering granular control over content delivery through configurable filters. Dockerized for easy deployment and scaling.


## Built for Efficiency:

 - Performance: Handles high request volumes with minimal latency, ensuring a smooth user experience.
  
 - Scalability: Adapts seamlessly to growing workloads, allowing you to scale your applications effectively.

# Getting Started:

Before diving in, ensure you have the following:

 - Make: Streamlines build processes.
    
 - Docker: Simplifies containerized deployments.

### 1. Clone the repository

```
git clone https://github.com/4kpros/cdn.git
cd go-cdn/
```
 - By default, the container uses `.env.example` as a sample environment file. You can directly update this file with your values. However, it would be better to integrate Vault to manage your secrets in your CI/CD pipeline. Or if you user Kubernetes.

### 2. Build and start the container

```
make docker-cdn
```
API documentation (using OpenAPI v3.1) at: http://localhost:23100/api/v1/docs
![DOC](https://github.com/user-attachments/assets/89afe1f7-a100-4a49-b492-c92f2717e8a1)


