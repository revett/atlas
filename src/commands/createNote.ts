import * as dayjs from "dayjs";
import * as vscode from "vscode";
import * as createNoteLib from "../lib/createNote";
import { openDocumentation } from "../lib/openDocumentation";
import { ToastType, toast } from "../lib/toast";

// Types of actions within the vscode.QuickPick.
enum QuickPickActionType {
  entity = "entity",
  meeting = "meeting",
  project = "project",
  scratch = "scratch",
  system = "system",
  topic = "topic",
  help = "help",
}

// Labels within the vscode.QuickPick for each action.
const quickPickActions: { [key in QuickPickActionType]: string } = {
  entity: "$(location) Entity",
  meeting: "$(organization) Meeting",
  project: "$(file-directory-create) Project",
  scratch: "$(flame) Scratch",
  system: "$(circuit-board) System",
  topic: "$(preview) Topic",
  help: "$(question) Help",
};

// Default note content for each action.
const noteDefaultContent: { [key in QuickPickActionType]: string } = {
  entity: `## TODO: Who? / What? / Where?

Note about TODO`,
  meeting: `# What?

Meeting about TODO with TODO.`,
  project: `## What?

New project about TODO`,
  scratch: `## What?

Random note about TODO`,
  system: `## What?

A system for TODO.

## Checklist

- TODO
- TODO
- TODO`,
  topic: `## What?

Note about TODO`,
  help: "",
};

// Convert the raw input from the user in to a dot case + kebab case.
const convertInputToFilename = (str: string): string => {
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

export const createNote = (): vscode.Disposable => {
  return vscode.commands.registerCommand("atlas.createNote", async () => {
    const quickPick = vscode.window.createQuickPick();

    // See: https://code.visualstudio.com/api/references/icons-in-labels#icon-listing
    quickPick.items = [
      {
        label: quickPickActions.entity,
        detail: "Note about a specific thing (e.g. location, person, company).",
      },
      {
        label: quickPickActions.meeting,
        detail: "Minutes and actions about a meeting (e.g. 1:1 with Alex).",
      },
      {
        label: quickPickActions.project,
        detail:
          "Short-term project with a specific goal and deadline (e.g. renovate bathroom).",
      },
      {
        label: quickPickActions.scratch,
        detail: "Temporary throwaway note to capture a passing thought.",
      },
      {
        label: quickPickActions.system,
        detail:
          "Checklist that you use to consistently complete a task (e.g. monthly budget).",
      },
      {
        label: quickPickActions.topic,
        detail: "Notes about a topic that you are interested in (e.g. coffee).",
      },
      {
        label: quickPickActions.help,
        description:
          "Not sure which type to use? Open the documentation in GitHub.",
      },
    ];
    quickPick.placeholder = "Select note type...";

    quickPick.onDidChangeSelection(async (selection) => {
      if (selection.length !== 1) {
        toast(
          ToastType.error,
          `Expected a single note type to selected, got ${selection.length} instead.`
        );
      }

      // Map the selected QuickPick label to an action type.
      const actionType = Object.entries(quickPickActions).find(
        ([key, value]) => value === selection[0].label
      )?.[0] as QuickPickActionType;

      // Variable to control what happens next.
      let shouldCreateNote = false;
      let requireFilenameInput = false;
      let filenameInputPrefix = "";
      let filename = "";

      // Logic for each of the action types.
      switch (actionType) {
        case QuickPickActionType.entity:
          shouldCreateNote = true;
          requireFilenameInput = true;
          // Important to use [] as this string will be used within dayjs.Format().
          filenameInputPrefix = `[${QuickPickActionType.entity}]`;
          break;

        case QuickPickActionType.meeting:
          shouldCreateNote = true;
          requireFilenameInput = true;
          filenameInputPrefix = `[${QuickPickActionType.meeting}].YYYY.MM.DD.HHmm`;
          break;

        case QuickPickActionType.project:
          shouldCreateNote = true;
          requireFilenameInput = true;
          filenameInputPrefix = `[${QuickPickActionType.meeting}].YYYY.[Q]Q`;
          break;

        case QuickPickActionType.scratch:
          shouldCreateNote = true;
          filename = "[scratch].YYYY.MM.DD.HHmmss.[md]";
          break;

        case QuickPickActionType.system:
          shouldCreateNote = true;
          requireFilenameInput = true;
          filenameInputPrefix = `[${QuickPickActionType.system}]`;
          break;

        case QuickPickActionType.topic:
          shouldCreateNote = true;
          requireFilenameInput = true;
          filenameInputPrefix = `[${QuickPickActionType.topic}]`;
          break;

        case QuickPickActionType.help:
          openDocumentation();
          break;

        default:
          // This is some TS magic to ensure that all enum values of quickPickActionType have logic
          // in the switch statement.
          const assertNever: never = actionType;
          toast(
            ToastType.error,
            "Unknown note type selected, unable to create new note."
          );
          break;
      }

      // Show an InputBox for action types that have dynamic filenames.
      if (requireFilenameInput && filenameInputPrefix !== "") {
        const inputBox = vscode.window.createInputBox();
        inputBox.placeholder = "Enter note title...";
        inputBox.show();

        await new Promise((resolve) => {
          inputBox.onDidChangeValue((e) => {
            if (e === "") {
              inputBox.prompt = "";
            }

            if (e !== "") {
              const ts = dayjs();
              inputBox.prompt = ts.format(
                `${filenameInputPrefix}.[${convertInputToFilename(e)}.md]`
              );
            }
          });

          inputBox.onDidAccept(() => {
            resolve(inputBox.value);
            inputBox.dispose();
          });
        });

        filename = `${filenameInputPrefix}.[${convertInputToFilename(
          inputBox.value
        )}.md]`;
      }

      // Create the note if instructed to.
      if (shouldCreateNote && filename !== "") {
        createNoteLib.createNote(filename, noteDefaultContent[actionType]);
      }
    });
    quickPick.onDidHide(() => quickPick.dispose());
    quickPick.show();
  });
};
