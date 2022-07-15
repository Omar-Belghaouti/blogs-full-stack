import { Blog } from "models";

export const testBlogs: Blog[] = [
  {
    id: 1,
    title: "Blog 1",
    body: "Body 1",
    author: 1,
    comments: [1, 2],
  },
  {
    id: 2,
    title: "Blog 2",
    body: "Body 2",
    author: 1,
    comments: [3, 4],
  },
];
