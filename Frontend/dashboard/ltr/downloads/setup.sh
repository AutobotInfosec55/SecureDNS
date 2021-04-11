wget 'http://101.53.147.32/dashboard/ltr/downloads/client.py'

virtualenv -p python2 niceEnv

source niceEnv/bin/activate

pip install tkintertable
pip install requests
pip install scapy


sudo python client.py
