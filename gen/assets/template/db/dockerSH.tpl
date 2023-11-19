#!/bin/bash
docker build -t abelce/{{lowerCase .Name}}_db:1.0.0 .
docker push abelce/{{lowerCase .Name}}_db:1.0.0