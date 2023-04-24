import * as dayjs from "dayjs";
import * as vscode from "vscode";
import { generateID } from "../lib/id";
import { createNote } from "../lib/createNote";

const generateContent = (ts: dayjs.Dayjs): string => {
  return `---
id: ${generateID()}
created: ${ts.format("ddd, DD MMM YYYY HH:mm:ss z")}
---

## What?

Note about TODO
`;
};

const convertInput = (str: string): string => {
  // Removes all the punctuations except hyphens.
  const cleanStr = str.replace(/[^\w\s-]|_/g, "").toLowerCase();

  // Replaces all whitespace with a period.
  const periodStr = cleanStr.replace(/\s/g, ".");

  // Ensures that there is never more than one period or hyphen.
  const noDoubleChar = periodStr.replace(/(\.)\1+|(-)\1+/g, "$1$2");

  // Removes hyphens and periods from the start or end of the string.
  const noHyphenPeriodStartEnd = noDoubleChar.replace(/(^[-.]*)|([-]*)$/g, "");

  // If the end of the string is a period, remove it.
  const finalStr = noHyphenPeriodStartEnd.replace(/\.$/g, "");

  return finalStr;
};

export const createTopicNote = (): vscode.Disposable => {
  return vscode.commands.registerCommand("atlas.createTopicNote", async () => {
    const inputBox = vscode.window.createInputBox();
    inputBox.title = "Enter note title";
    inputBox.show();

    await new Promise((resolve) => {
      inputBox.onDidChangeValue((e) => {
        if (e === "") {
          inputBox.prompt = "";
        }

        if (e !== "") {
          inputBox.prompt = `topic.${convertInput(e)}.md`;
        }
      });

      inputBox.onDidAccept(() => {
        resolve(inputBox.value);
        inputBox.dispose();
      });
    });

    const now = dayjs();
    await createNote(
      `topic.${convertInput(inputBox.value)}.md`,
      generateContent(now)
    );
  });
};
