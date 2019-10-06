import { Component, OnInit, Input, EventEmitter, Output, ChangeDetectorRef } from '@angular/core';
import { SwipeDirection } from 'src/app/models/swipeDirection';

@Component({
  selector: 'app-decision-card',
  templateUrl: './decision-card.component.html',
  styleUrls: ['./decision-card.component.scss']
})
export class DecisionCardComponent implements OnInit {

  @Input() hot: boolean;

  @Input() title: string;

  @Input() text: string;

  @Output() decided = new EventEmitter<SwipeDirection>();

  constructor(private cdRef: ChangeDetectorRef) { }

  ngOnInit() {
  }

  onSwipeLeft(evt) {
    console.log("Swiped left")
    this.decided.emit(SwipeDirection.Hot)
  }

  onSwipeRight(evt) {
    console.log("Swiped right")
    this.decided.emit(SwipeDirection.Not)
  }

}
