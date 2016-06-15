#!/bin/sh
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get -y install htop

wget -O dp.0.2.0 https://github.com/charleswong/darkpassenger/releases/download/v0.2.0/dp-x86_64-linux
cp ./dp.0.2.0 /usr/bin/
sudo chmod +x /usr/bin/dp.0.2.0
sudo echo '#!/bin/sh -e' > /etc/rc.local
sudo echo 'nohup /usr/bin/dp.0.2.0 -listen=DP_SERVER -regiservice=:1091 -secret=$DP_PASSWORD -clientmode=false -crypto=aes-128-cfb &' >> /etc/rc.local

sudo reboot
