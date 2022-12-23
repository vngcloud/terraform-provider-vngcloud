#cloud-config
users:
  - name: stackops
    shell: /bin/bash
    sudo: ALL=(ALL) NOPASSWD:ALL
    lock_passwd: false
ssh_pwauth: True
runcmd:
  - [ bash, -c, "echo -e 'Aa@123123\nAa@123123' | passwd stackops"]
  - passwd --expire stackops
  - sed -i "s/^#Port 22/Port 234/" /etc/ssh/sshd_config
  - sed -i "s/^Port 22/Port 234/" /etc/ssh/sshd_config
  - systemctl restart sshd
  - firewall-cmd --remove-service ssh
  - firewall-cmd --remove-service ssh --permanent
  - firewall-cmd --add-port 234/tcp
  - firewall-cmd --add-port 234/tcp --permanent
power_state:
  mode: reboot
write_files:
  - content: |
      #!/bin/bash
      iptables -I INPUT 1 -p tcp --dport 234 -j ACCEPT
    path: /var/lib/cloud/scripts/per-boot/iptable.sh
    owner: root:root
    permissions: '700'