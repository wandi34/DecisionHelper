import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DecisionComponent } from './decision/decision-stack/decision-stack.component';



@NgModule({
  declarations: [DecisionComponent],
  imports: [
    CommonModule
  ],
  exports: [
    DecisionComponent
  ]
})
export class DecisionModule { }
