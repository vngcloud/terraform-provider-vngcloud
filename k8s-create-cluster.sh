#!/bin/bash
IPS=$1

chmod 600 ~/.ssh/id_rsa
for ip in $IPS
do
    cat << EOF >> ~/.ssh/config
Host $ip
    HostName $ip
    User stackops
    Port 234
    IdentityFile /home/stackops/.ssh/id_rsa
EOF
done
cat ~/.ssh/config

sudo apt-get update && sudo DEBIAN_FRONTEND=noninteractive apt-get install python3-pip -y
git clone https://github.com/kubernetes-sigs/kubespray
cd kubespray/
sudo pip3 install -r requirements.txt
cp -rfp inventory/sample inventory/mycluster
declare -a IPS=($IPS)
CONFIG_FILE=inventory/mycluster/hosts.yaml python3 contrib/inventory_builder/inventory.py ${IPS[@]}
cat inventory/mycluster/hosts.yaml
ansible-playbook -i inventory/mycluster/hosts.yaml  --become --become-user=root cluster.yml