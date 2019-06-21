#!/usr/bin/env bash
rm -r realestate-peer
rm -r regulator-peer
rm -r trader-peer
rm -r uploads

mongo localhost:27017/estate --eval "db.dropDatabase()"
