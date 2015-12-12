#!/bin/bash
# -----------------------------------------------------------------
# NEO4J API Requests
# -----------------------------------------------------------------

# Change Password
curl -X GET -H "Content-Type: application/json" -H "Accept: application/json; charset=UTF-8" -H "Authorization: Basic bmVvNGo6c2VjcmV0" -d '{"username" : "neo4j", "password" : "neo4j", "password_change" : "http://localhost:7474/user/neo4j/password","password_change_required" : true}' 'http://localhost:7474/user/neo4j'
