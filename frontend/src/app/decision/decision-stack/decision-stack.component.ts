import {
  Component,
  OnInit,
  ViewChild,
  ElementRef,
  ChangeDetectorRef,
  AfterViewInit
} from "@angular/core";
import { Decision } from "../../models/decision";
import { DecisionService } from "../decision.service";
import { Topic } from "../../models/topic";
import { SwipeDirection } from "src/app/models/swipeDirection";

@Component({
  selector: "dt-decision-stack",
  templateUrl: "./decision-stack.component.html",
  styleUrls: ["./decision-stack.component.scss"]
})
export class DecisionComponent implements OnInit, AfterViewInit {
  @ViewChild("decisionStack", { static: false }) elementView: ElementRef;
  headerHeight: number;
  decisionList: Decision[];

  constructor(
    private decisionService: DecisionService,
    private cdRef: ChangeDetectorRef
  ) {}

  ngOnInit() {
    this.decisionList = this.decisionService.loadDecisionsForTopic(
      new Topic("Tribe Name")
    );
  }

  ngAfterViewInit() {
    this.headerHeight = this.elementView.nativeElement.offsetHeight * -1;
    this.cdRef.detectChanges();
  }

  onDecided(evt: SwipeDirection) {
    if (evt == SwipeDirection.Hot) {
      this.decisionList[0].hot = true;
      this.decisionService.voteForDecision(this.decisionList[0]);
      setTimeout(() => {
        this.decisionList.shift();
      }, 500);
    } else {
      this.decisionList[0].hot = false;
      this.decisionService.voteForDecision(this.decisionList[0]);
      setTimeout(() => {
        this.decisionList.shift();
      }, 500);
    }
  }
}
