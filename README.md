# Cryptography

Программа, разделяет секретную строку в hex формате длиной 256 bit (например, приватный ключ ECDSA secp256k1) на N частей по схеме Шамира и восстанавливает его при предъявлении любых T частей.
To compile and run: 

```shell
// разделение секрета:
go run . -split
// восстановление секрета:
go run . -recover
// запуск теста:
go run . -test
```
<b>Входные данные</b> - mode split (stdin):
1 строчка - Приватный ключ (P_KEY)
2 строчка - Два числа N, T, где 2 < T <= N < 100

<b>Выходные данные</b> - mode split (stdout):
N строк, в каждой строке содержится кусочек разделенного секрета(в виде string) + точка X

```shell
stdin:
a11b0a4e1a132305652ee7a8eb7848f6ad5ea381e3ce20a2c086a2e388230811
Enter N and T...
7 3

stdout:
keys:
7 01ca8ea9755e2f6df3dc0d81d8b4a088991a78eac21487df913a5ee63f2f8361d2
3 017c4f94c8f11c074ca77ccdc3471310d35f37a57cfe77d5b70858270eab51d6bd
2 4646eaccb2d58ae0bc2c8d183f2c2d67e1cd2ca4d8f89fa7b0bd6a421d3673b2
1 96c1040299ad701a77533e9434ccb73729fc671c9f4d3d020452eeb304185239
5 01495dd9e5aec1d543a27a3b555d858a30813eeb9356538785aed61e68c72eb885
4 9f9555be3d5fbd7551c00b78d38896e48d6e6ee0347bc4e985ff021cfec1a6d8
6 01ad064b22b9b1bac325e962caa88d856cb0426a33f6579068404b6a742cc4a183
```

<b>Входные данные</b> - mode recover (stdin):
T или более строк с кусочками секрета + точка X (в каждой строке кусок секрета в том же формате что и вывод программы в режиме split)

<b>Выходные данные</b> - mode recover (stdout):
Приватный ключ (P_KEY), в таком же виде что и перед разделением

```shell
stdin:
Enter sharing points...
3 b705a84cb04f3f29cecf5607e061c3c74366ebee9087c37ee42ccf13b7811a7e
5 01066f886b321569ad42d7a42c706f0f9136dd8c78fe4f0a411d49627a140e2665
8 121fb4ab959b391ec18e25acef3112d3f0d274492d43ae886266deff2282cd6c

output:
a11b0a4e1a132305652ee7a8eb7848f6ad5ea381e3ce20a2c086a2e388230811 <---- This is your key     
```