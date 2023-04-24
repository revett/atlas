import * as dayjs from "dayjs";
import * as vscode from "vscode";
import { generateID } from "../lib/id";
import { createNote } from "../lib/createNote";

const generateContent = (now: dayjs.Dayjs): string => {
  return `---
id: ${generateID()}
created: ${now.format("ddd, DD MMM YYYY HH:mm:ss z")}
---

## What?

Random note about TODO
`;
};

export const createScratchNote = (): vscode.Disposable => {
  return vscode.commands.registerCommand(
    "atlas.createScratchNote",
    async () => {
      const now = dayjs();
      await createNote(
        `scratch.${now.format("YYYY.MM.DD.HHmmss")}.md`,
        generateContent(now)
      );
    }
  );
};
