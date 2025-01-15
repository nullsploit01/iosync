export const parseSearchParamToNumber = (str: string | string[]): number => {
  if (!str || typeof str !== 'string') throw new Error('Invalid search param')
  return parseInt(str)
}
