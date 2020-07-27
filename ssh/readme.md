### dropbear


```bash

dropbear -h
Dropbear server v2015.67 https://matt.ucc.asn.au/dropbear/dropbear.html
Usage: dropbear [options]
-b bannerfile	Display the contents of bannerfile before user login
		(default: none)
-r keyfile  Specify hostkeys (repeatable)
		defaults: 
		dss /etc/dropbear/dropbear_dss_host_key
		rsa /etc/dropbear/dropbear_rsa_host_key
-R		Create hostkeys as required
-F		Don't fork into background
-E		Log to stderr rather than syslog
-w		Disallow root logins
-s		Disable password logins
-g		Disable password logins for root
-B		Allow blank password logins #!!!
-j		Disable local port forwarding
-k		Disable remote port forwarding
-a		Allow connections to forwarded ports from any host
-p [address:]port
		Listen on specified tcp port (and optionally address),
		up to 10 can be specified
		(default port is 22 if none specified)
-P PidFile	Create pid file PidFile
		(default /var/run/dropbear.pid)
-W <receive_window_buffer> (default 24576, larger may be faster, max 1MB)
-K <keepalive>  (0 is never, default 0, in seconds)
-I <idle_timeout>  (0 is never, default 0, in seconds)
-V    Version

```


```bash

cd /etc/dropbear
rm *
for i in rsa dss ecdsa ed25519 do
    dropbearkey -t ${i}  -f dropbear_${i}_host_key  1>>authorized_keys
done
sed -i '/:/d'  authorized_keys
chmod 600 *

mkdir openssh
for i in `ls *key` 
do  
    dropbearconvert dropbear openssh $i openssh/$i 
done


#dropbear -p 33 -B
#ssh root@192.168.0.106 -i dropbear_rsa_host_key -p 33
#ssh root@192.168.0.106 -p 33

#/etc/init.d/dropbear enable
#/etc/init.d/dropbear start

################################################################################
#vi /etc/init.d/dropbear1
#!/bin/sh /etc/rc.common

START=50
STOP=50
USE_PROCD=1
PROG=/usr/sbin/dropbear
NAME=dropbear1
PIDCOUNT=0

start_service() {
	/bin/busybox passwd root -d
	procd_open_instance
	procd_set_param command $PROG -B -r /etc/dropbear/dropbear_rsa_host_key
	procd_close_instance
}

shutdown() {
    stop
    /usr/bin/killall dropbear
}

################################################################################



chmod 777 /etc/init.d/dropbear1
/etc/init.d/dropbear1 start
ps |grep drop
13245 root      2268 S    /usr/sbin/dropbear -B -r /etc/dropbear/dropbear_rsa_host_key

ssh root@192.168.0.106 -i dropbear_rsa_host_key

```


### openssh



```bash

wget https://archive.openwrt.org/chaos_calmer/15.05/arm64/generic/packages/packages/openssh-keygen_7.1p2-1_arm64.ipk
tar xzvf openssh-keygen_7.1p2-1_arm64.ipk 
tar xzvf data.tar.gz -C /

wget https://archive.openwrt.org/chaos_calmer/15.05/arm64/generic/packages/packages/openssh-server_7.1p2-1_arm64.ipk
#opkg install openssh-server_7.1p2-1_arm64.ipk
tar xzvf openssh-server_7.1p2-1_arm64.ipk 
./debian-binary
./data.tar.gz
./control.tar.gz

tar xzvf data.tar.gz -C /

/usr/sbin/sshd
/etc/ssh/sshd_config
/etc/init.d/sshd


/usr/sbin/sshd
Privilege separation user sshd does not exist

在/etc/passwd 中加入： 
sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin
mkdir -p /var/empty/sshd



cd /etc/ssh/
for type in rsa dsa ecdsa ed25519
do 
  key=/etc/ssh/ssh_host_${type}_key
  echo $type $key
  ssh-keygen -N '' -t $type -f $key 2>&- >&-
done

chmod 600 /etc/ssh/*_key
/etc/ssh/sshd_config


Port 2222
AddressFamily any
ListenAddress 0.0.0.0
ListenAddress ::

HostKey /etc/ssh/ssh_host_rsa_key
HostKey /etc/ssh/ssh_host_dsa_key
HostKey /etc/ssh/ssh_host_ecdsa_key
HostKey /etc/ssh/ssh_host_ed25519_key

PasswordAuthentication yes
PermitEmptyPasswords yes


mkdir ~/.ssh
cat /etc/ssh/*.pub > ~/.ssh/authorized_keys
/etc/init.d/sshd enable
/etc/init.d/sshd start
#/usr/sbin/sshd

#copy ssh_host_rsa_key to pc
ssh 192.168.0.106 -i ssh_host_rsa_key -p 2222
```

