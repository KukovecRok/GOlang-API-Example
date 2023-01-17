# GOlang API

## Use case example
* Register on http://localhost:8000/api/v1/todo/login/register 
* .json: {
  "username":"username",
  "password":"password"
  } 
* Login: http://localhost:8000/api/v1/todo/login returns a JWT token
* GET request, add token, send to http://localhost:8000/api/v1/todo/opravilo/ 

### Primer uporabe
- Registracija na: http://localhost:8000/api/v1/todo/login/register 
- Json: {
  "username":"username",
  "password":"password"
  }
- Login: http://localhost:8000/api/v1/todo/login vrne JWT token
- GET request, priloži token na http://localhost:8000/api/v1/todo/opravilo/

## License

Copyright © 2023 [Tatookie](https://github.com/KukovecRok). <br /> 
This project is MIT licensed.

## Disclaimer

This framework is provided as-is, and there are no guarantees that it fits your purposes or that it is bug-free. Use it at your own risk!