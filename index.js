export function countWords(text = '') {
  const words = text.trim().match(/\p{L}[\p{L}\p{N}'’-]*/gu)
  return words ? words.length : 0
}

export function readingTime(text, PPM = 200) {
  if (PPM < 80) throw new Error('ppm deveria ser >= 80')
  const minutes = countWords(text) / PPM
  return Number(minutes.toFixed(2))
}