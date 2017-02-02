/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.10.dev12
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package evesi

/* 200 ok object */
type GetUniverseCategoriesCategoryIdOk struct {

	/* category_id integer */
	CategoryId int32 `json:"category_id,omitempty"`

	/* groups array */
	Groups []int32 `json:"groups,omitempty"`

	/* name string */
	Name string `json:"name,omitempty"`

	/* published boolean */
	Published bool `json:"published,omitempty"`
}
