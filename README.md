# DevOps with Kubernetes

My answers for the [DevOps with Kubernetes course](https://devopswithkubernetes.com/)

## Status

Completed 4/5 parts

## DBaaS vs DIY

### DBaaS

#### Pros

- **Easy to deploy**: Most DBaaS services have a web interface that allows you to configure and deploy a database in the click of a button. N oneed to manage Kubernetes configurations.
- **Easy to manage**: Like deploying, you can also easily manage your database using the web interface. Need additional storage? Need multiple replicas? Need a backup? Depending on the DBaaS provider, almost everything can be done using a simple web interface. Depending on provider you also have automated updates, security fixed and high availability.
- **Reliable**: The database is always up and running. If something goes wrong, you can just blame the service provider and they will fix it.
- **Secure**: The database is likely to be more secure than a DIY solution since the service provider usually knows what they are doing.

#### Cons

- **Cost**: The DBaaS is likely expensive than a DIY solution since on top of hosting, the service provider gets a fee.
- **Less customizable**: The DBaaS is not as customizable as a DIY solution since you don't have access to all the configurations.
- **Vendor lock-in**: Although the DBaaS provider usually knows what they are doing, you are entirely at their mercy. A very crucial part of your application is now dependent on a third party. For large companies this can be a turn off.

### DIY

#### Pros

- **Reliable**: The database is always up and running. If something goes wrong, you can just blame yourself and fix it.
- **Cheap**: The database is likely to be cheaper since you are only paying for hosting. In Google Cloud, you only play for GKE resources and Persistent Disks.
- **Anything is possible**: Since you own the database solution, you can do anything you want with it. You can even create your own web interface to manage it. You also have alot more conrol over resource allocation and optimization.

#### Cons

- **More work**: Configuring, deploying and managing a DIY database is usually much more work than a DBaaS solution.
- **More complicated**: Configuring, deploying and managing a DIY database is also usually much more complicated than a DBaaS solution, requiring more or less technical knowledge to do it. Scaling and backup management also require alot more overhead to implement than in a DBaaS.

## AKS vs DigitalOcean Kubernetes

### AKS (Azure Kubernetes Service)

#### Pros

- **Good for enterprises**: Robust security and many compliance options
- **Vast ecosystem**: The Azure Ecosystem contains alot of different products and services that integrate seamlessly with AKS
- **Flexible**: Support Windows and Linux nodes, advanced scaling, supports hybrid clouds

#### Cons 

- **Complexity**: The Azure ecosystem requires some knowledge and the initial setup might be harder than of DigitalOcean
- **Cost**: Pricing can be high depending on the type of services used

### DigitalOcean Kubernetes

#### Pros

- **Simple**: DigitalOcean is known for its simple and easy UI
- **Cheap**: DigitalOcean has affordable and straightforward billing making it better for teams with budget constraints
- **Quick Setup**: Creation of cluster is very fast and easy, has automated repairs and overall requires very little oversight


#### Cons

- **Not good for enterprises**: Lacks security and compliance tools
- **Doesn't scale**: DigitalOcean Kubernetes doesn't have as many scaling options as AKS and is not suitable for multi-cloud
- **Less support**: DigitalOcean has support, but it is not as extensive as AKS and may not meet the needs of a large enterprise

### Which is better?

If one had to choose, overall Azure Kubernetes Service (AKS) is better as it is much more flexible and powerful. However, DigitalOcean is still a very good option depending on circumstances.
