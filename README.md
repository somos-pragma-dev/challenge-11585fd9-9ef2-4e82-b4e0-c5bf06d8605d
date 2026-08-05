# Diseño y Despliegue de Microservicios en Go

En un entorno de banca digital, necesitas diseñar y desplegar microservicios en Go que interactúen mediante gRPC. Los servicios deben manejar transacciones financieras y comunicarse con un sistema de gestión de usuarios. El sistema debe soportar un throughput de 10 000 transacciones por segundo con una latencia máxima de 50ms. Debes asegurar la idempotencia de las transacciones mediante una clave única generada por el sistema. El fallo del servicio de usuarios debe ser manejado de manera que no afecte la disponibilidad del servicio de transacciones.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Microservices |
| **Nivel** | junior-l1 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Un IDE o editor de código.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Verifica que el proyecto arranca sin errores.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Definición de Servicios

**Objetivo:** Identificar y definir los microservicios necesarios para el dominio de transacciones financieras.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Enumera los microservicios requeridos.
- Define las operaciones que cada microservicio debe soportar.
- Establece las claves de idempotencia para las transacciones.

**Entregable:** Diagrama de servicios y operaciones, junto con la definición de claves de idempotencia.

<details>
<summary>Pistas de conocimiento</summary>

- Considera la modularidad y la separación de preocupaciones.
- Piensa en cómo los servicios se comunicarán entre sí.

</details>

### Fase 2: Implementación de Microservicios

**Objetivo:** Implementar los microservicios definidos utilizando Go y gRPC.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Crea los microservicios en Go.
- Implementa las operaciones definidas en la fase anterior.
- Asegura la idempotencia de las transacciones utilizando las claves definidas.

**Entregable:** Código fuente de los microservicios implementados en Go con gRPC.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza patrones de diseño adecuados para la implementación.
- Asegúrate de que los servicios sean idempotentes.

</details>

### Fase 3: Despliegue en Kubernetes

**Objetivo:** Desplegar los microservicios en un entorno de Kubernetes.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Configura un cluster de Kubernetes.
- Despliega los microservicios en el cluster.
- Verifica que los servicios se comuniquen correctamente y cumplan con los requisitos de rendimiento y latencia.

**Entregable:** Cluster de Kubernetes con los microservicios desplegados y funcionando.

<details>
<summary>Pistas de conocimiento</summary>

- Utiliza herramientas de monitoreo para verificar el rendimiento y la latencia.
- Asegúrate de que los servicios sean escalables y resilientes.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué son los microservicios y por qué se utilizan en este dominio?
- **paraQueSirve**: ¿Para qué sirven las claves de idempotencia en las transacciones financieras?
- **comoSeUsa**: ¿Cómo se implementan y despliegan los microservicios en un entorno de Kubernetes?
- **erroresComunes**: ¿Cuáles son los errores comunes al implementar microservicios en Go y cómo se pueden evitar?
- **queDecisionesImplica**: ¿Qué decisiones implica el diseño y despliegue de microservicios en un entorno de banca digital?

## Criterios de Evaluacion

- Definir correctamente los microservicios y sus operaciones.
- Implementar los microservicios en Go con gRPC de manera idempotente.
- Desplegar los microservicios en un entorno de Kubernetes y verificar su funcionamiento.

---

*Reto generado automaticamente por Challenge Generator - Pragma*
