Test de threading entre python y go

En mi resultados la diferencia es aproximadamente x2 ganando go
Estaría bueno medir los recursos de cpu que utlizan los procesos
que es importante en ambientes productivos

Para compilar go usar:

	go build -buildmode=c-shared -o shfp.so shfp.go

test hechos con python 3.6