FROM continuumio/miniconda3 AS poincare_build
SHELL ["/bin/bash", "-c"]
RUN apt install gcc g++ -y
RUN git clone https://github.com/facebookresearch/poincare-embeddings.git
WORKDIR /poincare-embeddings/ 
RUN pip install  nltk sklearn torch==1.0.0 pandas h5py cython tqdm
RUN pip install --upgrade cython
RUN python setup.py build_ext --inplace
COPY ./teste.csv .
CMD ["python3", "embed.py", "-dim", "2","-lr", "2", "-epochs", "200","-negs", "50","-burnin", "20", "-dampening", "0.75","-ndproc",  "4","-manifold", "poincare","-dset", "teste.csv","-checkpoint", "ptbr.pth","-batchsize", "80","-eval_each", "10","-fresh","-sparse","-train_threads", "2"]
