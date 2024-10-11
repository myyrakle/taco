mkdir -p /app
cp out/taco /app/taco
cp taco.service /etc/systemd/system/taco.service
systemctl restart taco.service
systemctl daemon-reload
systemctl status taco.service