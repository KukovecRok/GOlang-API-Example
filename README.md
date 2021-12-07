### Primer uporabe
- Registracija na: http://localhost:8000/api/v1/todo/login/register 
- Json: {
  "username":"username",
  "password":"password"
  }
- Login: http://localhost:8000/api/v1/todo/login vrne token
- GET request, priloži token na http://localhost:8000/api/v1/todo/opravilo/ 