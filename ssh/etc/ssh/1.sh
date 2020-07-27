for type in rsa dsa ecdsa ed25519
do
  key=ssh_host_${type}_key
  echo $type $key
  ssh-keygen -N '' -t $type -f $key 2>&- >&-
done

