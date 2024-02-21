
// 根据分类返回目录
export const getDir =  (categories : string[]): string => {
  // 如果分类包含lcr，返回lcr目录
  if (categories.includes('LCR')) {
    return 'lcr'
  }
  return "leetcode"
}

