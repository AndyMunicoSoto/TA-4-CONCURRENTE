mtype { COMPUTE, FINISH };

#define TESTS 5
#define DATA_SIZE 10
chan result_channel = [2];

typedef DataPoint {
    int Size;
    int Price;
}

DataPoint data[DATA_SIZE]; // Arreglo para almacenar los datos

proctype linear_regression() {

    int sumX = 0, sumY = 0, sumXY = 0, sumXSquare = 0, slope = 0, intercept = 0;

    int i;
    
    // Cálculo de la regresión lineal

    i = 0;
    do
    :: i < DATA_SIZE ->
        sumX = sumX + data[i].Size;
        sumY = sumY + data[i].Price;
        sumXY = sumXY + data[i].Size * data[i].Price;
        sumXSquare = sumXSquare + data[i].Size * data[i].Size;
        i = i + 1;
    :: else ->
        break; // Salir del bucle cuando i >= DATA_SIZE
    od;

    int n = DATA_SIZE;
    slope = (n * sumXY - sumX * sumY) / (n * sumXSquare - sumX * sumX) ;
    intercept = (sumY - slope * sumX) / n;
    
    // Enviar los resultados a través del canal
    result_channel!slope;
    result_channel!intercept;
}

proctype handle_results() {
    int slope, intercept;
    
    // Recibir los resultados de los procesos de regresión lineal
    do
    :: result_channel??slope; result_channel??intercept -> 
       // Almacenar los resultados
       results[_pid * 2] = slope;
       results[_pid * 2 + 1] = intercept;
    od
    
    // Indicar que todos los procesos han terminado
    all_finished = 1;
}

init {
    int i;

    all_finished = 0;
    // Inicializar los datos y ejecutar los procesos de regresión lineal
    i = 0;

    do
    :: i < TESTS ->
        run linear_regression();
        i = i + 1;
    :: else ->
        break; // Salir del bucle cuando i >= TESTS
    od;

    // Ejecutar el proceso para manejar los resultados
    run handle_results();
    // Esperar a que todos los procesos terminen
    do
    :: all_finished == 1 -> break;
    od
}
