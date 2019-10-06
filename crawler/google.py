from google.cloud import translate
client = translate.Client()
print(client)
project_id = 'red-inspr'
location = 'global'
 
parent = client.location_path(project_id, location)

# edges = set()
# for synset in tqdm(wn.all_synsets(pos='n')):
#     # write the transitive closure of all hypernyms of a synset to file
#     for hyper in synset.closure(lambda s: s.hypernyms()):
#         word = client.translate_text(
#             parent=parent,
#             contents=[synset.name()[:synset.name().find('.')]],
#             mime_type='text/plain',  # mime types: text/plain, text/html
#             source_language_code='en-US',
#             target_language_code='pt').translations[0]
#         hypert = client.translate_text(
#             parent=parent,
#             contents=[synset.name()[:hyper.name().find('.')]],
#             mime_type='text/plain',  # mime types: text/plain, text/html
#             source_language_code='en-US',
#             target_language_code='pt').translations[0]
#         print(word)
#         print(hypert)
#         edges.add((word, hyper))

#     # also write transitive closure for all instances of a synset
#     for instance in synset.instance_hyponyms():
#         for hyper in instance.closure(lambda s: s.instance_hypernyms()):
#             edges.add((instance.name(), hyper.name()))
#             for h in hyper.closure(lambda s: s.hypernyms()):
#                 edges.add((instance.name(), h.name()))

