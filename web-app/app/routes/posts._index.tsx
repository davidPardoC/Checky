import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { getPosts } from "~/models/post.server";

export const loader = async () => {
  return json({
    posts: await getPosts()
  });
};

export default function Posts() {
    const data = useLoaderData<typeof loader>()
    return (
      <main>
        <h1>Posts</h1>
      </main>
    );
  }