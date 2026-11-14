import random
import time

frases = [
    "Las papas están observándote 🥔",
    "Error 404: motivación no encontrada",
    "Un pato acaba de aprobar tu código 🦆",
    "Python dice: hola humano",
    "Este mensaje se autodestruirá... mentira"
]

colores = ["🔴", "🟢", "🔵", "🟡", "🟣"]

print("Iniciando programa!!!\n")
time.sleep(1)

for i in range(random.randint(3, 7)):
    print(random.choice(colores), random.choice(frases))
    time.sleep(random.uniform(0.3, 1.2))

if random.choice([True, False]):
    print("\n🎉 ¡Felicidades! No pasó nada malo :)")
else:
    print("\n🤖 El programa decidió terminar dramáticamente?")

print("\nFin del caos!!!!!")
