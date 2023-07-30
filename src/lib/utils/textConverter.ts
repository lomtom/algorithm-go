import { slug } from "github-slugger";
import { marked } from "marked";
import {html} from "mdast-util-to-markdown/lib/handle/html";

marked.use({
  mangle: false,
  headerIds: false,
});

// slugify
export const slugify = (content: string): string => {
  return slug(content);
};

// markdownify
export const markdownify = (content: string, div?: boolean): string => {
  return div ? marked.parse(content) : marked.parseInline(content);
};

// humanize
export const difficulty = (content: string): string => {
  if (content == "简单") {
    return "🌟"
  }else if (content == "中等") {
    return "🌟🌟🌟"
  }else if (content == "困难") {
    return "🌟🌟🌟🌟🌟"
  }
  return humanize(content)
}

// humanize
export const humanize = (content: string): string => {
  if (content == "" || content == undefined){
    return ""
  }
  return content
    .replace(/^[\s_]+|[\s_]+$/g, "")
    .replace(/[_\s]+/g, " ")
    .replace(/^[a-z]/, function (m) {
      return m.toUpperCase();
    });
};

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
