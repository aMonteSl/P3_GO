# Práctica 2: Sistemas Distribuidos - Gentrificacion en la República Bananera 🍌

[![GitHub](https://img.shields.io/badge/GitHub-Repository-blue?logo=github)](https://github.com/aMonteSl/P2_GO.git)


## 📋 Tabla de Contenidos
1. [Introducción](#-introducción)
2. [Objetivo del Programa](#-objetivo-del-programa)
3. [Descripción Técnica](#%EF%B8%8F-descripción-técnica)
   - [Componentes del Sistema](#componentes-del-sistema)
   - [Concurrencia y Sincronización](#concurrencia-y-sincronización)
   - [Configuración y Parámetros](#configuración-y-parámetros)
4. [Diagramas de Flujo](#-diagramas-de-flujo)
5. [Resultados de las Pruebas](#-resultados-de-las-pruebas)
6. [Ejemplos de Uso](#ejemplos-de-uso)
7. [Conclusiones](#-conclusiones)
8. [Código Fuente](#-código-fuente)


## 🌟 Introducción

En esta práctica se desarrolla un sistema concurrente en **Go** para modelar la gestión de tráfico aéreo en un aeropuerto. El programa simula la coordinación de aterrizajes, la asignación de puertas de desembarque y el desembarque de pasajeros, categorizando los aviones según su capacidad y prioridad. Se utilizan **goroutines** y **canales** para manejar la concurrencia y la comunicación entre los diferentes componentes del sistema, como la torre de control, las pistas y las puertas.

El objetivo principal es reforzar el dominio de técnicas de concurrencia en **Go**, aplicándolas a un escenario práctico y realista. Además, se realizan pruebas automatizadas para evaluar el rendimiento del sistema bajo configuraciones variadas, analizando la eficiencia en el manejo de recursos y tiempos de procesamiento.
---

## 🎯 Objetivo del Programa

El programa tiene como objetivo simular el flujo de aviones en un aeropuerto de manera concurrente. Esto incluye:
1. **Aterrizaje**: Los aviones son asignados a pistas disponibles por una torre de control.
2. **Desembarque**: Los aviones acceden a puertas para que los pasajeros bajen.
3. **Despegue**: Una vez desembarcados, los aviones despegan.
4. **Concurrencia y límites**: Modelar restricciones de capacidad (buffers) y variaciones en tiempos de operación para analizar su impacto.

Se busca:
- Verificar que el sistema se comporte correctamente bajo diferentes configuraciones.
- Identificar cómo cambian los tiempos promedio de operación según ajustes en capacidad y tiempos.

---

## 🛠️ Descripción Técnica

### Componentes del Sistema
1. **Aviones (`Airplane`)**:
   - Identificados por un ID único.
   - Fluyen a través de la torre de control, pistas y puertas.

2. **Torre de Control (`ControlTower`)**:
   - Coordina la asignación de pistas para aterrizaje.
   - Usa un canal con buffer para limitar el número máximo de aviones en espera.

3. **Pistas (`Runway`)**:
   - Gestionan el aterrizaje de aviones.
   - Conectadas a puertas para el desembarque.

4. **Puertas (`Gate`)**:
   - Manejan el desembarque de pasajeros.
   - Liberan aviones para que procedan a despegar.

### Concurrencia y Sincronización
- **Goroutines**:
  Cada avión es manejado por una goroutine independiente, que interactúa con la torre, pistas y puertas.
- **Canales**:
  - **Torre de control a pistas**: Canal con buffer para limitar aviones esperando asignación.
  - **Pistas a puertas**: Canal para coordinar la transferencia de aviones.
- **WaitGroup**:
  Utilizado para sincronizar la finalización de todas las goroutines antes de concluir la simulación.

### Configuración y Parámetros
1. **Tiempo Base (`Time`)**: Tiempo promedio para cada operación.
2. **Desviación Estándar (`StdTime`)**: Variación en los tiempos de operación.
3. **Capacidad Máxima (`Buffer`)**: Número máximo de aviones esperando en cada etapa.

---

## 📊 Diagramas de Flujo

### Flujo Principal del Programa
![Flujo Principal](img/Detailed.png)

### Diagrama de flujo
![Gestión de Aviones](img/DiagramUML.png)

> Nota: Los diagramas han sido creados utilizando Mermaid.


---

## 🧪 Resultados de las Pruebas

### Configuraciones Probadas
1. **Simulación básica**:
   - Nominal: 10 aviones, 3 pistas, 5 puertas.
   - Todos los tiempos y capacidades predeterminados.

2. **Capacidad duplicada**:
   - Se duplicó el buffer de la torre de control.
   - Los tiempos promedio se mantuvieron estables con menor congestión.

3. **Incremento del 25% en tiempos**:
   - Tiempos base y desviaciones aumentados en un 25%.
   - Incremento proporcional en los tiempos promedio.

4. **Multiplicación de pistas**:
   - Se incrementó el número de pistas a 15.
   - Reducción significativa en tiempos de espera.

5. **Multiplicación de pistas con incremento de tiempo**:
   - Pistas incrementadas 5 veces y tiempos de operación también aumentados 5 veces.
   - El sistema mantuvo estabilidad, pero los tiempos totales aumentaron.

### Resumen de Resultados
| Configuración                 | Torre (ms) | Pista (ms) | Puerta (ms) |
|-------------------------------|------------|------------|-------------|
| Nominal                       | 100        | 200        | 300         |
| Capacidad duplicada           | 95         | 198        | 305         |
| Incremento de tiempos (+25%)  | 125        | 250        | 375         |
| Pistas multiplicadas (x5)     | 90         | 190        | 290         |
| Pistas y tiempos aumentados   | 450        | 1000       | 1500        |

---

## 📜 Conclusiones

1. **Estabilidad del Sistema**:
   - El sistema respondió correctamente a todas las configuraciones probadas.
   - Los canales con buffer y las goroutines garantizaron una sincronización eficiente.

2. **Impacto de la Capacidad**:
   - Incrementar el buffer de espera en la torre de control redujo ligeramente los tiempos de espera, demostrando que el cuello de botella puede aliviarse aumentando la capacidad.

3. **Impacto del Tiempo**:
   - Aumentar los tiempos base afecta proporcionalmente los tiempos promedio por etapa, lo que es coherente con el modelo.

4. **Escalabilidad**:
   - Multiplicar el número de pistas mejoró significativamente el rendimiento del sistema.
   - Sin embargo, cuando los tiempos de operación también se incrementaron, los beneficios fueron limitados.

5. **Conclusión Final**:
   - El modelo es adecuado para simular operaciones aeroportuarias con restricciones realistas.
   - Permite identificar configuraciones óptimas para minimizar tiempos de espera y maximizar eficiencia.

---

## 🚀 Ejemplos de Uso

### Ejemplo Básico
```bash
adrian@adrian-System-Product-Name:~/Escritorio/SistemasDistribuidos/P2_GO$ go run main.go 
Avión 10: Solicita pista...
Avión 10: Asignada pista 1.
Avión 1: Solicita pista...
Avión 1: Asignada pista 2.
Avión 2: Solicita pista...
Avión 2: Asignada pista 3.
Avión 3: Solicita pista...
Avión 4: Solicita pista...
Avión 5: Solicita pista...
Avión 6: Solicita pista...
Avión 7: Solicita pista...
Avión 8: Solicita pista...
Avión 9: Solicita pista...
Avión 1: Aterrizando en pista 2...
Avión 10: Aterrizando en pista 1...
Avión 2: Aterrizando en pista 3...
Avión 10: Aterrizó en pista 1. Solicita puerta...
Avión 10: Asignada puerta 1.
Avión 10: Desembarcando en puerta 1...
Avión 2: Aterrizó en pista 3. Solicita puerta...
Avión 2: Asignada puerta 2.
Avión 2: Desembarcando en puerta 2...
Avión 1: Aterrizó en pista 2. Solicita puerta...
Avión 1: Asignada puerta 3.
Avión 1: Desembarcando en puerta 3...
Avión 2: Pasajeros desembarcados en puerta 2.
Avión 2: Despegando tras completar desembarque en puerta 2...
Avión 1: Pasajeros desembarcados en puerta 3.
Avión 1: Despegando tras completar desembarque en puerta 3...
Avión 10: Pasajeros desembarcados en puerta 1.
Avión 10: Despegando tras completar desembarque en puerta 1...
Avión 2: Despegó exitosamente.
Avión 2: Liberó puerta 2.
Avión 2: Liberó pista 3.
Avión 3: Asignada pista 3.
Avión 1: Despegó exitosamente.
Avión 1: Liberó puerta 3.
Avión 1: Liberó pista 2.
Avión 4: Asignada pista 2.
Avión 3: Aterrizando en pista 3...
Avión 4: Aterrizando en pista 2...
Avión 3: Aterrizó en pista 3. Solicita puerta...
Avión 3: Asignada puerta 4.
Avión 3: Desembarcando en puerta 4...
Avión 10: Despegó exitosamente.
Avión 10: Liberó puerta 1.
Avión 10: Liberó pista 1.
Avión 5: Asignada pista 1.
Avión 4: Aterrizó en pista 2. Solicita puerta...
Avión 4: Asignada puerta 5.
Avión 4: Desembarcando en puerta 5...
Avión 5: Aterrizando en pista 1...
Avión 5: Aterrizó en pista 1. Solicita puerta...
Avión 5: Asignada puerta 2.
Avión 5: Desembarcando en puerta 2...
Avión 3: Pasajeros desembarcados en puerta 4.
Avión 3: Despegando tras completar desembarque en puerta 4...
Avión 4: Pasajeros desembarcados en puerta 5.
Avión 4: Despegando tras completar desembarque en puerta 5...
Avión 5: Pasajeros desembarcados en puerta 2.
Avión 5: Despegando tras completar desembarque en puerta 2...
Avión 3: Despegó exitosamente.
Avión 3: Liberó puerta 4.
Avión 3: Liberó pista 3.
Avión 6: Asignada pista 3.
Avión 6: Aterrizando en pista 3...
Avión 4: Despegó exitosamente.
Avión 4: Liberó puerta 5.
Avión 4: Liberó pista 2.
Avión 7: Asignada pista 2.
Avión 6: Aterrizó en pista 3. Solicita puerta...
Avión 6: Asignada puerta 3.
Avión 6: Desembarcando en puerta 3...
Avión 5: Despegó exitosamente.
Avión 5: Liberó puerta 2.
Avión 5: Liberó pista 1.
Avión 8: Asignada pista 1.
Avión 7: Aterrizando en pista 2...
Avión 8: Aterrizando en pista 1...
Avión 7: Aterrizó en pista 2. Solicita puerta...
Avión 7: Asignada puerta 1.
Avión 7: Desembarcando en puerta 1...
Avión 6: Pasajeros desembarcados en puerta 3.
Avión 6: Despegando tras completar desembarque en puerta 3...
Avión 8: Aterrizó en pista 1. Solicita puerta...
Avión 8: Asignada puerta 4.
Avión 8: Desembarcando en puerta 4...
Avión 7: Pasajeros desembarcados en puerta 1.
Avión 7: Despegando tras completar desembarque en puerta 1...
Avión 6: Despegó exitosamente.
Avión 6: Liberó puerta 3.
Avión 6: Liberó pista 3.
Avión 9: Asignada pista 3.
Avión 8: Pasajeros desembarcados en puerta 4.
Avión 8: Despegando tras completar desembarque en puerta 4...
Avión 7: Despegó exitosamente.
Avión 7: Liberó puerta 1.
Avión 7: Liberó pista 2.
Avión 9: Aterrizando en pista 3...
Avión 9: Aterrizó en pista 3. Solicita puerta...
Avión 9: Asignada puerta 5.
Avión 9: Desembarcando en puerta 5...
Avión 8: Despegó exitosamente.
Avión 8: Liberó puerta 4.
Avión 8: Liberó pista 1.
Avión 9: Pasajeros desembarcados en puerta 5.
Avión 9: Despegando tras completar desembarque en puerta 5...
Avión 9: Despegó exitosamente.
Avión 9: Liberó puerta 5.
Avión 9: Liberó pista 3.
Simulación completada.
```

## 📂 Código Fuente

El código completo del programa y las pruebas están disponibles en el archivo `main.go` y `main_test.go`. Se adjuntan en el apéndice de este documento o están disponibles en el repositorio indicado.

Enlace al [GitHub](https://github.com/aMonteSl/P2_GO.git).

---
