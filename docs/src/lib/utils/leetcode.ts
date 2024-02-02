
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

async function getQuestionOfToday(): Promise<QuestionOfTodayResponse|undefined> {
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

export { getQuestionOfToday};
