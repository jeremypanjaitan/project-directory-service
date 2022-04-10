import { API_CATEGORY, API_PROJECT, axios } from "../../config";
import { useFetchCall } from "../../hooks";

export const useCategoryData = () => {
		const categoryData = async () => {
				try {
						const res = await axios.get(API_CATEGORY);
						const data = res.data.data;
						return data;
				} catch (err) {
						throw err;
				}
		};
		return useFetchCall(categoryData);
};

export const useCategoryDataByID = () => {
		const categoryData = async (id) => {
				try {
						const res = await axios.get(API_CATEGORY + "/" + id);
						const data = res.data.data;
						return data;
				} catch (err) {
						throw err;
				}
		};
		return useFetchCall(categoryData);
};

export const useProjectData = () => {
		const projectData = async (id) => {
				try {
						const res = await axios.get(API_PROJECT + "/" + id);
						const data = res.data.data;
						return data;
				} catch (err) {
						throw err;
				}
		};
		return useFetchCall(projectData);
};

export const useEditProject = () => {
		const addProject = async ({id, values}) => {
				try {
						const url = API_PROJECT + "/" + id + "/"

						const res = await axios.put(url, {title: values.title, picture: values.picture,
								story: values.story, description: values.description, categoryId : values.categoryId});
						const data = res.data;
						return data;
				} catch (err) {
						throw err;
				}
		};
		return useFetchCall(addProject);
};
