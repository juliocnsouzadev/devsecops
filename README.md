# Devsecops

## Installing Talisman
Pre commit checks for security issues
```bash
curl https://thoughtworks.github.io/talisman/install.sh > /tmp/install-talisman.sh
chmod +x /tmp/install-talisman.sh
cd /root/kubernetes-devops-security
bash /tmp/install-talisman.sh pre-commit
```