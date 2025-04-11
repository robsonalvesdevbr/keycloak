import base64
import hashlib
import os

def generate_code_verifier():
    return base64.urlsafe_b64encode(os.urandom(32)).decode('utf-8').rstrip('=')

def generate_code_challenge(code_verifier):
    hashed = hashlib.sha256(code_verifier.encode('utf-8')).digest()
    return base64.urlsafe_b64encode(hashed).decode('utf-8').rstrip('=')

code_verifier = generate_code_verifier()
code_challenge = generate_code_challenge(code_verifier)

print(f"Code Verifier: {code_verifier}")
print(f"Code Challenge (S256): {code_challenge}")