#!/bin/sh
sudo apt-get update
sudo apt-get -y upgrade
sudo apt-get -y install htop
sudo apt-get -y install squid
sudo apt-get -y install supervisor

sudo service squid3 stop
#sudo service squid stop

sudo update-rc.d squid3 disable
#sudo update-rc.d squid disable

curl -L https://raw.githubusercontent.com/CharlesWong/darkpassenger/master/vps/squid.conf > tmp.conf
sudo cat tmp.conf >> /etc/squid3/squid.conf
#sudo cat tmp.conf >> /etc/squid/squid.conf
#sudo service squid restart

wget -O dp https://github.com/CharlesWong/darkpassenger/raw/master/vps/dp
cp ./dp /usr/bin/
sudo chmod +x /usr/bin/dp

curl -L https://raw.githubusercontent.com/CharlesWong/darkpassenger/master/vps/supervisord.conf > tmp.conf
sudo cat tmp.conf > /etc/supervisor/conf.d/dp.conf

#sudo echo '#!/bin/sh -e' > /etc/rc.local
#sudo echo '/usr/bin/dp -listen=:9156 -secret=qwertyuiop123 -crypto=rc4 -backend=127.0.0.1:8888 -clientmode=false &' >> /etc/rc.local

sudo reboot