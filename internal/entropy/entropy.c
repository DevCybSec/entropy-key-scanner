#include <math.h>

/**
 * Calcula la Entropía de Shannon de un buffer de datos.
 * Optimizada para no realizar asignaciones en el heap.
 */
double calculate_entropy(const unsigned char* data, int size) {
    if (size == 0) return 0.0;

    // Array de frecuencias en el stack.
    // Un byte tiene 256 valores posibles (0-255).
    unsigned int counts[256] = {0};

    // Paso 1: Contar frecuencias (O(n))
    for (int i = 0; i < size; ++i) {
        counts[data[i]]++;
    }

    // Paso 2: Calcular entropía (O(256) -> Constante)
    double entropy = 0.0;
    double log2_inv = 1.4426950408889634; // 1/log(2) para optimizar la división

    for (int i = 0; i < 256; ++i) {
        if (counts[i] > 0) {
            double p = (double)counts[i] / size;
            entropy -= p * log(p) * log2_inv;
        }
    }

    return entropy;
}