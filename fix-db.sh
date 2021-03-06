
mysql -f --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e "STOP SLAVE"
mysql -f --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e "RESET SLAVE"

mysqldump --dump-slave --master-data --gtid --password=$CONFIGURATION_DATABASE_PASSWORD --user=root --host=$CONFIGURATION_DATABASE_REPLICATION_SETUP_HOST --all-databases > /tmp/dump.sql

mysql -f --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD < /tmp/dump.sql

mysql -f --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e "FLUSH PRIVILEGES"

#Set Master
mysql -f --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e "CHANGE MASTER TO master_host='plume.byu.edu', master_port=3306, master_user='$CONFIGURATION_DATABASE_USERNAME', master_password='$CONFIGURATION_DATABASE_PASSWORD', master_use_gtid=slave_pos;"

#START SLAVE;
mysql --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e 'START SLAVE';
mysql --user=$CONFIGURATION_DATABASE_USERNAME --password=$CONFIGURATION_DATABASE_PASSWORD -e 'SHOW SLAVE STATUS';
