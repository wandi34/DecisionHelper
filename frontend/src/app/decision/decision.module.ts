import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DecisionComponent } from './decision-stack/decision-stack.component'
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { CustomMaterialModule } from '../custom-material.module';
import { DecisionCardComponent } from './decision-card/decision-card.component';

@NgModule({
  declarations: [DecisionComponent, DecisionCardComponent],
  imports: [
    CommonModule,
    CustomMaterialModule,
    BrowserAnimationsModule
  ],
  exports: [
    DecisionComponent
  ]
})
export class DecisionModule { }
