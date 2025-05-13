"use client";

import React, { useState, useEffect } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Separator } from "@/components/ui/separator";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { ArrowDownIcon } from "@radix-ui/react-icons";
import { Button } from "@/components/ui/button";
import { ArrowRightIcon } from "@radix-ui/react-icons";
import { convertMarkdownToHtml } from "@/lib/markdown";
import { getHashKey } from "@/lib/utils";
import { apiFetch } from "@/lib/apifetch";
interface Question {
  number: string | null;
  text: string | null;
  options: string[];
  discussionLink: string | null;
  answerLink: string | null;
  otherLink: string | null;
}

const ranges = [
  "2025-04-08-12%3A00%3A00-professional-cloud-architect%201-50.md", 
  "2025-04-09-12%3A00%3A00-professional-cloud-architect%2051-100.md", 
  "2025-04-10-12%3A00%3A00-professional-cloud-architect%20101-150.md", 
  "2025-04-11-12%3A00%3A00-professional-cloud-architect%20151-200.md", 
  "2025-04-11-12%3A00%3A00-professional-cloud-architect%20201-250.md", 
  "2025-04-11-12%3A00%3A00-professional-cloud-architect%20251-279.md"];
function getMarkdownUrl(index: number) {
  return "https://raw.githubusercontent.com/Axpz/xMinima/refs/heads/master/_posts/" + ranges[index];
}

function getDiscussionUrl(index: string) {
  return "https://raw.githubusercontent.com/Axpz/xMinima/refs/heads/master/assets/gcp/discussion/" + index + ".md";
}

function getAnswerUrl(index: string) {
  return "https://raw.githubusercontent.com/Axpz/xMinima/refs/heads/master/assets/gcp/discussion/" + index + ".ans.html";
} 

function getOtherUrl(index: string) {
  return "https://raw.githubusercontent.com/Axpz/xMinima/refs/heads/master/assets/gcp/discussion/" + index + ".vs.html";
}

async function fetchRemoteContent(url: string) {
  const cacheKey = getHashKey(url);
  const cached = localStorage.getItem(cacheKey);

  if (cached && cached.length > 100) {
    // 简单校验长度，避免空值或垃圾数据
    return cached;
  }

  try {
    console.log("fetching", url);
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const text = await response.text();

    if (text && text.length > 100) {
      localStorage.setItem(cacheKey, text);
    }

    return text;
  } catch (error) {
    console.error("Could not fetch markdown file:", error);
    return "";
  }
}

function parseQuestions(markdownContent: string | null): Question[] {
  if (!markdownContent) {
    return [];
  }

  const questions: Question[] = [];
  const lines = markdownContent.split("\n");
  let currentQuestion: Question = {
    number: null,
    text: null,
    options: [],
    discussionLink: null,
    answerLink: null,
    otherLink: null,
  };
  let readingQuestionText = false;

  for (const line of lines) {
    const numberMatch = line.match(/^#\s+(\d+)\s*$/);
    if (numberMatch) {
      if (currentQuestion.number) {
        questions.push(currentQuestion);
        currentQuestion = {
          number: null,
          text: null,
          options: [],
          discussionLink: null,
          answerLink: null,
          otherLink: null,
        };
        readingQuestionText = false;
      }
      currentQuestion.number = numberMatch[1];
      readingQuestionText = true;
      continue;
    }

    if (
      readingQuestionText &&
      !line.startsWith("[") 
      // &&
      // !line.match(/^[A-G]\.\s/
      // )
    ) {
      currentQuestion.text = currentQuestion.text
        ? `${currentQuestion.text}\n${line}`
        : line;
      continue;
    }

    // const optionMatch = line.match(/^([A-G])\.\s(.*)$/);
    // if (optionMatch) {
    //   currentQuestion.options.push(line);
    //   continue;
    // }

    const discussionMatch = line.match(/^\[Discussion]\((.*)\)\s*$/);
    if (discussionMatch) {
      currentQuestion.discussionLink = discussionMatch[1];
      continue;
    }

    const answerMatch = line.match(/^\[.nswer.*]\((.*)\)\s*$/);
    if (answerMatch) {
      currentQuestion.answerLink = answerMatch[1];
      continue;
    }

    const otherMatch = line.match(/^\[.*]\((.*)\)\s*$/);
    if (otherMatch) {
      currentQuestion.otherLink = otherMatch[1];
      continue;
    }
  }

  if (currentQuestion.number) {
    questions.push(currentQuestion);
  }

  return questions;
}

export default function QuestionsPage({
  params,
}: {
  params: Promise<{ index: string }>;
}) {
  const { index } = React.use(params);
  const pageIndex = parseInt(index, 10);
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [expandedDiscussions, setExpandedDiscussions] = useState<Record<string, boolean>>({});
  const [discussions, setDiscussions] = useState<Record<string, string>>({});
  const [expandedAnswers, setExpandedAnswers] = React.useState<Record<string, boolean>>({});
  const [answers, setAnswers] = React.useState<Record<string, string>>({});
  const [expandedOthers, setExpandedOthers] = React.useState<Record<string, boolean>>({});
  const [others, setOthers] = React.useState<Record<string, string>>({});

  const toggleDiscussionExpand = async (index: string) => {
    setExpandedDiscussions(prev => ({ ...prev, [index]: !prev[index] }));
    if (!discussions[index]) {
      const discussion = await fetchRemoteContent(getDiscussionUrl(index))
      const html = await convertMarkdownToHtml(discussion)
      setDiscussions(prev => ({ ...prev, [index]: html }));
    }
  };

  const toggleAnswerExpand = async (index: string) => {
    setExpandedAnswers(prev => ({ ...prev, [index]: !prev[index] }));
    if (!answers[index]) {
      const ans = await fetchRemoteContent(getAnswerUrl(index))
      const html = await convertMarkdownToHtml(ans)
      setAnswers(prev => ({ ...prev, [index]: html }));
    }
  };

  const toggleOtherExpand = async (index: string) => {
    setExpandedOthers(prev => ({ ...prev, [index]: !prev[index] }));
    if (!others[index]) {
      const html = await fetchRemoteContent(getOtherUrl(index))
      setOthers(prev => ({ ...prev, [index]: html }));
    }
  };

  useEffect(() => {
    const loadContent = async () => {
      try {
        const markdown = await fetchRemoteContent(getMarkdownUrl(pageIndex));
        if (markdown) {
          const parsedQuestions = parseQuestions(markdown);
          setQuestions(parsedQuestions);
        }
      } catch (err: any) {
        setError(err.message || "Failed to load questions.");
      } finally {
        setLoading(false);
      }
    };

    loadContent();
  }, []);

  if (loading) {
    return <div className="p-4">Loading questions...</div>;
  }

  if (error) {
    return <div className="p-4 text-red-500">{error}</div>;
  }

  return (
    <div className="flex flex-col gap-4 max-w-4xl mx-auto items-start">
      {" "}
      {/* 移除 p-6 */}
      <h1 className="text-2xl font-bold text-gray-900 dark:text-gray-100">
        Professional Cloud Architect Questions
      </h1>
      {questions.map((question) => (
        <Card key={question.number} className="w-full shadow-md">
          <CardHeader>
            <CardTitle className="text-xl text-left">
              Question {question.number}
            </CardTitle>
          </CardHeader>
          <CardContent className="text-lg space-y-2 text-gray-800 dark:text-gray-200 items-start">
            {question.text && (
              <p className="whitespace-pre-line text-left">{question.text}</p>
            )}
            {question.options.map((option, index) => (
              <p key={index} className="ml-4 text-left">
                {option}
              </p>
            ))}
            <div className="mt-4 flex flex-col gap-0 items-start test-left">
              {question.discussionLink && (
                <Button
                  variant="ghost"
                  onClick={() => toggleDiscussionExpand(question.number || "")}
                  className="inline-flex items-center justify-start text-blue-500 hover:text-blue-700 hover:bg-blue-50 dark:hover:bg-blue-900/20 w-48 mt-2"
                >
                  {expandedDiscussions[question.number || ""] ? (
                    <ArrowDownIcon
                      className="mr-2 h-4 w-4" aria-hidden="true"
                    />
                  ) : (
                    <ArrowRightIcon
                      className="mr-2 h-4 w-4" aria-hidden="true"
                    />
                  )}
                  <span className="text-left">Discussion</span>
                </Button>
              )}

              {expandedDiscussions[question.number || ""] && (
                <div className="border-l-4 border-gray-300 pl-4 ml-2">
                  <div className="prose mx-auto text-left">
                    <div dangerouslySetInnerHTML= {{ __html: discussions[question.number || ""] || "Loading..." }} />
                  </div>
                </div>
              )}

              {question.answerLink && (
                <Button
                  variant="ghost"
                  onClick={() => toggleAnswerExpand(question.number || "")}
                  className="inline-flex items-center justify-start text-green-500 hover:text-green-700 hover:bg-green-50 dark:hover:bg-green-900/20 w-48 mt-2"
                >
                  {expandedAnswers[question.number || ""] ? (
                    <ArrowDownIcon
                      className="mr-2 h-4 w-4"
                      aria-hidden="true"
                    />
                  ) : (
                    <ArrowRightIcon
                      className="mr-2 h-4 w-4"
                      aria-hidden="true"
                    />
                  )}
                  <span className="text-left">Answer & Vote</span>
                </Button>
              )}
              {expandedAnswers[question.number || ""] && (
                <div className="prose mx-auto text-left">
                  <div dangerouslySetInnerHTML= {{ __html: answers[question.number || ""] || "Loading..." }} />
                </div>
              )}

              {question.otherLink && (
                <Button
                  variant="ghost"
                  onClick={() => toggleOtherExpand(question.number || "")}
                  className="inline-flex items-center justify-start text-purple-500 hover:text-purple-700 hover:bg-purple-50 dark:hover:bg-purple-900/20 w-48 mt-2"
                >
                  {expandedOthers[question.number || ""] ? (
                    <ArrowDownIcon
                      className="mr-2 h-4 w-4"
                      aria-hidden="true"
                    />
                  ) : (
                    <ArrowRightIcon
                      className="mr-2 h-4 w-4"
                      aria-hidden="true"
                    />
                  )}
                  <span className="text-left">Other</span>  
                </Button>
              )}
              {expandedOthers[question.number || ""] && (
                <div className="border-l-4 border-gray-300 pl-4 ml-2">  
                  <div className="prose mx-auto text-left">
                    <div dangerouslySetInnerHTML= {{ __html: others[question.number || ""] || "Loading..." }} />
                  </div>
                </div>
              )}

            </div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}
