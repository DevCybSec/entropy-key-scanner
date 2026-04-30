#ifndef ENTROPY_H
#define ENTROPY_H

/**
 * Calcula la Entropía de Shannon de un buffer de datos.
 * @param data Puntero al array de bytes.
 * @param size Tamaño del buffer.
 * @return Valor de entropía entre 0.0 y 8.0.
 */
double calculate_entropy(const unsigned char* data, int size);

#endif