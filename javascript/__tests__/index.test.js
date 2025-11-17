import { countWords, readingTime } from '../index.js'

const sample = 'Simples como voar. Chega ser lacrimejante. Giropops Strigus Girus'

test('contagem de palavras', () => {
  expect(countWords(sample)).toBe(9)
})

test('tempo de leitura padrão', () => {
  expect(readingTime(sample)).toBe(0.04)
})

test('erro quando PPM < 80', () => {
  expect(() => readingTime(sample, 50)).toThrow()
})
