import * as dayjs from "dayjs";
import * as vscode from "vscode";
import { openDocumentation } from "./commands/openDocumentation";
import { createScratchNote } from "./commands/createScratchNote";
import { createTopicNote } from "./commands/createTopicNote";
import { createNote } from "./commands/createNote";
var advancedFormatPlugin = require("dayjs/plugin/advancedFormat");
var timezonePlugin = require("dayjs/plugin/timezone");

export function activate(context: vscode.ExtensionContext) {
  dayjs.extend(advancedFormatPlugin);
  dayjs.extend(timezonePlugin);

  context.subscriptions.push(
    createNote(),
    createScratchNote(),
    createTopicNote(),
    openDocumentation()
  );
}

export function deactivate() {}
