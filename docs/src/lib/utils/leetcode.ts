interface QuestionOfTodayResponse {
  todayRecord: {
    date: string;
    userStatus: string;
    question: {
      questionId: number;
      frontendQuestionId: number;
      difficulty: string;
      title: string;
      titleCn: string;
      titleSlug: string;
      paidOnly: boolean;
      freqBar: string;
      isFavor: boolean;
      acRate: string;
      status: string;
      solutionNum: number;
      hasVideoSolution: boolean;
      topicTags: {
        name: string;
        nameTranslated: string;
        id: string;
      }[];
      extra: {
        topCompanyTags: {
          imgUrl: string;
          slug: string;
          numSubscribed: number;
        }[];
      };
    };
    lastSubmission: {
      id: string;
    };
  }[];
}

async function getQuestionOfToday(): Promise<QuestionOfTodayResponse | undefined> {
  return await fetch('https://leetcode.cn/graphql/', {
    method: 'POST',
    headers: {'Accept-Language': 'zh-CN,zh;q=0.9', 'Content-Type': 'application/json',},
    body: JSON.stringify({
      query: ` query questionOfToday { todayRecord { date userStatus question { questionId frontendQuestionId: questionFrontendId difficulty title titleCn: translatedTitle titleSlug paidOnly: isPaidOnly freqBar isFavor acRate status solutionNum hasVideoSolution topicTags { name nameTranslated: translatedName id } extra { topCompanyTags { imgUrl slug numSubscribed } } } lastSubmission { id } } } `,
      variables: {},
      operationName: 'questionOfToday',
    }),
  }).then(response => response.json())
    .then(data => {
      return data.data;
    })
    .catch(error => console.error('Error:', error))
}

// 判断array中的data.slug是否包含titleSlug，包含返回data.slug
export const getQuestionUrl = (titleSlug: string, blog_folder: string, array: any[]): string => {
  if (!titleSlug) {
    return '';
  }
  for (let i = 0; i < array.length; i++) {
    if (array[i].data.slug == titleSlug) {
      return blog_folder + "/" + array[i].data.slug;
    }
  }
  return "https://leetcode.cn/problems/" + titleSlug;
}


interface QuestionStatsResponse {
  data: {
    question: {
      stats: string;
    };
  };
}

interface QuestionStats {
  totalAccepted: string,
  totalSubmission: string,
  totalAcceptedRaw: string,
  totalSubmissionRaw: string,
  acRate: string,
}

async function getQuestionStatsInline(titleSlug: string): Promise<QuestionStats | void> {
  if (!titleSlug) {
    return;
  }
  return await fetch('https://leetcode.cn/graphql/', {
    method: 'POST',
    headers: {
      'Accept-Language': 'zh-CN,zh;q=0.9',
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: `
          query questionStats($titleSlug: String) {
            question(titleSlug: $titleSlug) {
              stats
            }
          }
        `,
      variables: {titleSlug},
      operationName: 'questionStats',
    }),
  }).then(response => response.json())
    .then(data => {
      return JSON.parse(data.data.question.stats)
    })
    .catch(error => console.error('Error:', error))
}


async function getQuestionStats(titleSlug: string): Promise<string | void> {
  if (!titleSlug) {
    return;
  }
  let data = await getQuestionStatsInline(titleSlug)
  if (data) {
    return `${data.acRate}`;
  }
  return ""
}

export {getQuestionOfToday, getQuestionStatsInline, getQuestionStats};
