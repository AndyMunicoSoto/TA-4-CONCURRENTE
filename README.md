# TP-CONCURRENTE

*Introducción*
El caso de predicción de precios de viviendas se resolverá mediante el uso del algoritmo de regresión lineal en función al tamaño de la casa, la cantidad de habitaciones, la antigüedad y la locación por sector. Esto se logrará a través de un programa realizado de manera concurrente y distribuida con el uso de canales en el lenguaje GO. Las pruebas se realizarán con 8 mil datos y se analizarán los resultados. También se usará la arquitectura cliente-servidor donde el cliente ingresará las variables para su predicción mediante la terminal. Posteriormente, se tendrá la conclusión del algoritmo sobre el caso de uso.

*Caso de uso*
Al comienzo del año se vio un incremento del 7% en ventas de viviendas en Lima Metropolitana con respecto al año pasado, esto se debe a que el cliente tiene una confianza positiva con las inmobiliarias (Gálvez, 2024). Y con el fin de mantener la fidelidad del cliente se requiere realizar un modelo de regresión lineal que prediga los precios de las viviendas. Las características a tomar para las predicciones son el tamaño de la vivienda (size), la cantidad de habitaciones (bedrooms), la antigüedad de la vivienda (age) y la locación por sector (location).

*Algoritmo*
Para la elaboración del algoritmo tenemos en cuenta el uso de la arquitectura cliente-servidor donde el cliente ingresa los parámetros para la predicción y el servidor procesa el modelo de regresión lineal y realiza la predicción devolviendo el precio al cliente.
El código se realizará en el algoritmo Go usando programación concurrente y distribuida con el uso de canales y el conjunto de datos de entrenamiento y prueba se establece en RAW.
