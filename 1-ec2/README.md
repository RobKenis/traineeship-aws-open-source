# Setting up a server

## Creating an EC2 instance

- Go to the **EC2** console
- Pick **launch instance**, it should be a bright orange button.
- Fill in the **Name** field, this will set the **name** tag on the server.
- Select an AMI, for this session, we will use **Amazon Linux 2 AMI (HVM), SSD Volume Type**
- Pick an instance type, it doesn't have to be big, **t2.micro** will work just fine
- Create a new key pair
  - enter a unique name
  - choose type **RSA** and format **.pem**
- Make sure a public IP will be assigned, this is enabled by default
- Add a security group with access to port **22** from **0.0.0.0/0**
- Launch the instance!

```shell
 aws ec2 run-instances --image-id ami-089950bc622d39ed8 --instance-type t2.nano --key-name rob --region eu-west-1
```

## Connecting to the EC2 instance

There are multiple ways to connect to a running EC2 instance, but I prefer **ssh**.

```shell
chmod 400 <key>.pem
ssh -i <key>.pem ec2-user@<public-ip>
```
