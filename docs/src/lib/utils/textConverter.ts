import {slug} from "github-slugger";
import {marked} from "marked";
import {html} from "mdast-util-to-markdown/lib/handle/html";

marked.use({
  mangle: false,
  headerIds: false,
});

// slugify
export const slugify = (content: string): string => {
  return slug(content);
};

// 百分比转换
export const percent = (content: string): string => {
  if (content == "" || content == undefined) {
    return ""
  }
  return (parseFloat(content) * 100).toFixed(1) + "%"
}

// markdownify
export const markdownify = (content: string, div?: boolean): string => {
  return div ? marked.parse(content) : marked.parseInline(content);
};

// humanize
export const difficulty = (content: string): string => {
  if (content == "简单" || content == "Easy") {
    return "🌟"
  } else if (content == "中等" || content == "Medium") {
    return "🌟🌟🌟"
  } else if (content == "困难" || content == "Hard") {
    return "🌟🌟🌟🌟🌟"
  }
  return humanize(content)
}
export const Undifficulty = (content: string): string => {
  if (content == "简单" || content == "Easy") {
    return "简单"
  } else if (content == "中等" || content == "Medium") {
    return "中等"
  } else if (content == "困难" || content == "Hard") {
    return "困难"
  }
  return humanize(content)
}

export const difficultyClass = (content: string): string => {
  if (content == "简单" || content == "Easy") {
    return "text-green-500 "
  } else if (content == "中等" || content == "Medium") {
    return "text-amber-500 "
  } else if (content == "困难" || content == "Hard") {
    return "text-red-500 "
  }
  return humanize(content)
}


export const difficultyClassBg = (content: string): string => {
  const a = "text-white px-3 py-1 rounded"
  if (content == "简单" || content == "Easy") {
    return "bg-green-400 " + a
  } else if (content == "中等" || content == "Medium") {
    return "bg-amber-300 " + a
  } else if (content == "困难" || content == "Hard") {
    return "bg-red-500 " + a
  }
  return humanize(content)
}


// humanize
export const humanize = (content: string): string => {
  if (content == "" || content == undefined) {
    return ""
  }
  return content
    .replace(/^[\s_]+|[\s_]+$/g, "")
    .replace(/[_\s]+/g, " ")
    .replace(/^[a-z]/, function (m) {
      return m.toUpperCase();
    });
};


export const categoryStar = (content: string): string => {
  if (content == "" || content == undefined) {
    return ""
  }
  return content + " " + difficulty(content)
}


// plainify
export const plainify = (content: string): string => {
  const filterBrackets = content.replace(/<\/?[^>]+(>|$)/gm, "");
  const filterSpaces = filterBrackets.replace(/[\r\n]\s*[\r\n]/gm, "");
  const stripHTML = htmlEntityDecoder(filterSpaces);
  return stripHTML;
};

// strip entities for plainify
const htmlEntityDecoder = (htmlWithEntities: string): string => {
  let entityList: { [key: string]: string } = {
    "&nbsp;": " ",
    "&lt;": "<",
    "&gt;": ">",
    "&amp;": "&",
    "&quot;": '"',
    "&#39;": "'",
  };
  let htmlWithoutEntities: string = htmlWithEntities.replace(
    /(&amp;|&lt;|&gt;|&quot;|&#39;)/g,
    (entity: string): string => {
      return entityList[entity];
    },
  );
  return htmlWithoutEntities;
};
