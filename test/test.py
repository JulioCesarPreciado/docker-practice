import requests
import jwt
import time

# 🔐 Configura tu secreto (igual al de tu .env)
JWT_SECRET = "secret"
JWT_ALGORITHM = "HS256"

# 🌐 URL del endpoint (ajústalo si es diferente)
BASE_URL = "http://localhost:8080"
CATALOG = "brands"  # o models, years, versions

# 🧾 Payload del token (puedes agregar claims si luego los validas)
payload = {
    "iat": int(time.time()),  # issued at
    "exp": int(time.time()) + 300  # expira en 5 minutos
}

# 🔑 Genera el token
token = jwt.encode(payload, JWT_SECRET, algorithm=JWT_ALGORITHM)

# 🚀 Hacer la petición POST
headers = {
    "Authorization": f"Bearer {token}"
}

response = requests.post(
                    f"{BASE_URL}/cache/refresh/{CATALOG}", headers=headers)

# 📋 Mostrar resultado
print("Status:", response.status_code)
print("Response:", response.text or "✅ Cache refreshed (No Content)")
