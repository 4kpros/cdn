# Customizable CDN with Go and Docker

- A flexible CDN built with Go, offering granular control over content delivery through configurable filters. Dockerized for easy deployment and scaling.


## Built for Efficiency:

 - Performance: Handles high request volumes with minimal latency, ensuring a smooth user experience.
  
 - Scalability: Adapts seamlessly to growing workloads, allowing you to scale your applications effectively.

# Getting Started:

Before diving in, ensure you have the following:

 - Make: Streamlines build processes.
    
 - Docker: Simplifies containerized deployments.

 - Rename .env.example to ```app.env```

### 1. Clone the repository

```
git clone https://github.com/4kpros/cdn.git
cd go-cdn/
```

### 2. Build and start the container

```
make docker-cdn
```
API documentation (using OpenAPI v3.1) at: http://localhost:23100/api/v1/docs
![DOC](https://github.com/user-attachments/assets/89afe1f7-a100-4a49-b492-c92f2717e8a1)


