# Getting started with Ansible in AWS

## Setting up Python
```shell
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
```

## Setting AWS credentials
```bash
export AWS_ACCESS_KEY_ID='<AWS_ACCESS_ID>'
export AWS_SECRET_ACCESS_KEY='<AWS_SECRET_KEY>'
```

## Setting up Ansible config
```
[defaults]
inventory = ./hosts/ec2.py
remote_user = ec2-user
private_key_file =  ../<key>.pem
deprecation_warnings=False
```

## Testing connection
```bash
ansible all -m ping
```

[Ansible: dynamic hosts](https://docs.ansible.com/ansible/latest/user_guide/intro_dynamic_inventory.html)

## Deploying an application
```shell
ansible-playbook application.yaml
```
