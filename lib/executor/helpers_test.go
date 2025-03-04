/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2020 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package executor

import "go.k6.io/k6/metrics"

func sumMetricValues(samples chan metrics.SampleContainer, metricName string) (sum float64) { //nolint:unparam
	for _, sc := range metrics.GetBufferedSamples(samples) {
		samples := sc.GetSamples()
		for _, s := range samples {
			if s.Metric.Name == metricName {
				sum += s.Value
			}
		}
	}
	return sum
}
