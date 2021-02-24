import { LitElement, html, css } from 'lit-element';
import { Theme } from '@furo/framework/src/theme';
// eslint-disable-next-line import/no-extraneous-dependencies

import { FBP } from '@furo/fbp';
// eslint-disable-next-line import/no-extraneous-dependencies
import '@furo/doc-helper';

// eslint-disable-next-line import/no-extraneous-dependencies
import '@furo/data-ui/src/furo-catalog.js';

import '@furo/data';
import '@furo/input';
import '@furo/app/src/furo-card';


/**
 * `demo-furo-data-chart`
 *
 * @customElement
 * @appliesMixin FBP
 */
class DemoFuroDataChartMixed extends FBP(LitElement) {

  /**
   * Themable Styles
   * @private
   * @return {CSSResult}
   */
  static get styles() {
    // language=CSS
    return Theme.getThemeForComponent(this.name) || css`
      :host {
        display: block;
        height: 100%;
        padding-right: var(--spacing);
      }

      :host([hidden]) {
        display: none;
      }
    `;
  }


  /**
   * @private
   * @returns {TemplateResult}
   */
  render() {
    // language=HTML
    return html`
      <furo-vertical-flex>

        <div>
          <h2>Demo ...</h2>
          <p>Describe your demo</p>
        </div>
        <furo-demo-snippet flex>

          <template>
            <style>
              furo-card {
                margin: 10px;

              }</style>

            <furo-vertical-scroller>

              <furo-card header-text="Multiple Data Sources in one chart">
              <furo-chart-display
                                  chart-type="line"
                                  title-text="XYZ - Stock Analysis"
                                  title-align="left"
                                  title-offset-x="70"
                                  no-data-text="Loading..."
                                  fixed-height="350"
                                  tooltip
                                  legend
                                  grid
                                  legend-align="left"
                                  legend-position="bottom"
                                  legend-offset-x="70"
                                  legend-offset-y="70"
                                  toolbar

              >


                <furo-data-chart-binder
                  ƒ-bind-data="--projectDAO(*.entities)"
                  data-field="data.cost_limit.units"
                  category-field="data.description"
                  legend-label="Cost"
                  chart-type="bar"
                  chart-color="#cd00fb"
                  chart-stroke-width="0"
                  chart-curve="smooth"
                  axis-label="Units"
                  axis-label-opposite
                  axis-ticks
                  axis-ticks-color="#FEA555"
                  axis-border
                  axis-border-color="#FEA555"
                  axis-label-color="#FEA232"
                  axis-tooltip
                ></furo-data-chart-binder>

                <furo-data-chart-binder
                  ƒ-bind-data="--projectDAO(*.entities)"
                  data-field="data.cost_limit.units"
                  category-field="data.description"
                  legend-label="Unit"
                  chart-stroke-width="4"
                  chart-curve="stepline"
                  axis-label="data.Unit._value"
                  axis-border
                  axis-tooltip
                ></furo-data-chart-binder>

                <furo-data-chart-binder
                  ƒ-bind-data="--projectDAO(*.entities)"
                  data-field="data.start.day"
                  category-field="data.description"
                  legend-label="Day"
                  axis-label="Day"
                  axis-ticks
                  axis-border
                  axis-tooltip
                  chart-marker-size="1"
                  x-chart-color="#cd00fb"
                  chart-stroke-width="4"
                  chart-curve="smooth"
                  y-axis-ticks
                  y-axis-border
                  y-axis-border-color="#008FFB"
                  y-axis-label-color="#008FFB"
                  tooltip
                  opposite
                ></furo-data-chart-binder>
              </furo-chart-display>

                <furo-button slot="action" label="load data" primary @-click="--btnListClicked"></furo-button>
              </furo-card>


              <furo-deep-link
                ƒ-trigger="--btnListClicked"
                service="ProjectService"
                @-hts-out="--hts"
              ></furo-deep-link>
              <furo-collection-agent
                service="ProjectService"
                ƒ-hts-in="--hts"
                list-on-hts-in
                @-response-hts-updated="--responseHts"
                @-response="--collectionResponse"
              >
              </furo-collection-agent>

              <furo-data-object
                type="project.ProjectCollection"
                ƒ-inject-raw="--collectionResponse"
                @-object-ready="--projectDAO"
              >
              </furo-data-object>

            </furo-vertical-scroller>
          </template>
        </furo-demo-snippet>
      </furo-vertical-flex>
    `;
  }
}

window.customElements.define('demo-furo-data-chart-mixed', DemoFuroDataChartMixed);
