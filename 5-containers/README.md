# Containers on AWS

## Building an application

```shell
docker build -t traineeship .
```

## Running the application

```shell
docker run -p 8080:8080 traineeship
```

## Pushing the application

### Setting up the registry
Deploy [registry.yaml](registry.yaml) using the CloudFormation console.
After creating the registry, navigate to **Elastic Container Registry** and click on your registry.
In the top right corner, you should see a button with **View push commands**

### Pushing your image
```shell
aws ecr get-login-password --region eu-west-1 | docker login --username AWS --password-stdin 811286505646.dkr.ecr.eu-west-1.amazonaws.com
```
```shell
docker tag traineeship:latest 811286505646.dkr.ecr.eu-west-1.amazonaws.com/<your-registry>:latest
```
```shell
docker push 811286505646.dkr.ecr.eu-west-1.amazonaws.com/<your-registry>:latest
```

### Verifying the image
```shell
docker run 811286505646.dkr.ecr.eu-west-1.amazonaws.com/<your-registry>:latest
```

## Deploying the application

### Creating an ECS cluster
Deploy [ecs_cluster.yaml](ecs_cluster.yaml) using the CloudFormation console.

### Creating an ECS service

### Exposing the application
